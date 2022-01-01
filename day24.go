package main

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Group struct {
	units  int64
	hp     int64
	atk    int64
	dmg    string
	ini    int64
	immune map[string]bool
	weak   map[string]bool
	id     int
	sel    bool
	tgt    *Group
	tdmg   int64
	army   *Army
	power  int64
}

type Army struct {
	name   string
	total  int64
	groups []*Group
	enemy  *Army
}

func (g *Group) info() string {
	return fmt.Sprintf("[%d] %d units, HP: %d DMG: %d %s INI: %d IMMUNE: %v WEAK: %v",
		g.id, g.units, g.hp, g.atk, g.dmg, g.ini, g.immune, g.weak)
}

func (g *Group) cmp(other *Group) bool {
	if g.power == other.power {
		return g.ini > other.ini
	}
	return g.power > other.power
}

func (g *Group) damage(other *Group) (ret int64) {
	ret = g.units * g.atk
	if other.immune[g.dmg] {
		ret = 0
	} else if other.weak[g.dmg] {
		ret *= 2
	}
	return ret
}

func (g *Group) target() {
	g.tdmg = 0
	g.tgt = nil
	for _, eg := range g.army.enemy.groups {
		if eg.sel || eg.units == 0 {
			continue
		}
		curd := g.damage(eg)
		if curd > g.tdmg || (curd > 0 && curd == g.tdmg && eg.cmp(g.tgt)) {
			g.tdmg = curd
			g.tgt = eg
		}
	}
	if g.tgt != nil {
		g.tgt.sel = true
	}
}

func (g *Group) attack() (kc int64) {
	if g.tgt != nil {
		kc = g.damage(g.tgt) / g.tgt.hp
		if kc > g.tgt.units {
			kc = g.tgt.units
		}
		g.tgt.units -= kc
		g.army.enemy.total -= kc
	}
	return
}

func (army Army) prepare(verbose bool, groups []*Group) []*Group {
	if verbose {
		fmt.Printf("%s: (total %d)\n", army.name, army.total)
	}
	for _, g := range army.groups {
		g.sel = false
		g.power = g.units * g.atk
		if verbose && g.units > 0 {
			fmt.Println(g.info())
		}
		groups = append(groups, g)
	}
	if verbose {
		fmt.Println()
	}
	return groups
}

func (army *Army) battle(enemy *Army, verbose bool) int64 {
	for {
		groups := army.prepare(verbose, make([]*Group, 0, 1000))
		groups = enemy.prepare(verbose, groups)
		if army.total == 0 || enemy.total == 0 {
			break
		}
		sort.Slice(groups, func(i, j int) bool { return groups[i].cmp(groups[j]) })
		for _, g := range groups {
			g.target()
			if g.tgt != nil && g.tdmg > 0 && verbose {
				fmt.Printf("%s group %d [units %d power %d] would deal defending group %d %d damage\n",
					g.army.name, g.id, g.units, g.power, g.tgt.id, g.tdmg)
			}
		}
		if verbose {
			fmt.Println()
		}
		sort.Slice(groups, func(i, j int) bool { return groups[i].ini > groups[j].ini })
		tkc := int64(0)
		for _, g := range groups {
			kc := g.attack()
			if kc > 0 && verbose {
				fmt.Printf("%s group %d [units %d power %d] attacks defending group %d killing %d\n",
					g.army.name, g.id, g.units, g.power, g.tgt.id, kc)
			}
			tkc += kc
		}
		if tkc == 0 {
			return -1
		}
		if verbose {
			fmt.Println()
		}
	}
	return army.total
}

func parseArmy(lines []string, modifier int64) (ret *Army) {
	r0 := regexp.MustCompile(`^(.*):`)
	r1 := regexp.MustCompile(`(\d+) units each with (\d+) hit points (\(.*?\) )?with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	r2 := regexp.MustCompile(`(immune|weak) to (.*)`)
	g := r0.FindStringSubmatch(lines[0])
	ret = &Army{g[1], 0, make([]*Group, 0, 100), nil}
	for _, l := range lines[1:] {
		g := r1.FindStringSubmatch(l)
		grp := &Group{}
		grp.units, _ = strconv.ParseInt(g[1], 10, 0)
		grp.hp, _ = strconv.ParseInt(g[2], 10, 0)
		grp.atk, _ = strconv.ParseInt(g[4], 10, 0)
		grp.atk += modifier
		grp.dmg = g[5]
		grp.ini, _ = strconv.ParseInt(g[6], 10, 0)
		grp.id = len(ret.groups) + 1
		grp.army = ret
		ret.groups = append(ret.groups, grp)
		ret.total += grp.units
		grp.weak = make(map[string]bool)
		grp.immune = make(map[string]bool)
		if len(g[3]) == 0 {
			continue
		}
		tg := g[3][1 : len(g[3])-2]
		for _, ss := range strings.Split(tg, "; ") {
			g2 := r2.FindStringSubmatch(ss)
			tl := strings.Split(g2[2], ", ")
			for _, s := range tl {
				sc := string([]byte(s))
				if g2[1] == "immune" {
					grp.immune[sc] = true
				} else if g2[1] == "weak" {
					grp.weak[sc] = true
				} else {
					fmt.Printf("Invalid ability kind: %s", g2[1])
					break
				}
			}
		}
	}
	return
}

func (x Aoc) Day24(scanner *bufio.Scanner) {
	chunks := make([][]string, 2)
	cid := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			cid++
			continue
		}
		chunks[cid] = append(chunks[cid], scanner.Text())
	}
	for power := int64(0); power < 1000000; power++ {
		immune := parseArmy(chunks[0], power)
		infection := parseArmy(chunks[1], 0)
		immune.enemy = infection
		infection.enemy = immune
		immune.battle(infection, power == 0)
		if power == 0 {
			fmt.Printf("Day1 Part 1: %d\n", infection.total+immune.total)
		}
		if immune.total > 0 && infection.total == 0 {
			fmt.Printf("Day1 Part 2: %d\n", immune.total)
			break
		} else if immune.total == 0 {
			fmt.Printf("Power %d: infection remaining %d\n", power, infection.total)
		} else {
			fmt.Printf("Power %d: tie %d/%d remaining\n", power, immune.total, infection.total)
		}
	}
}
