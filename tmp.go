package main

import (
	"fmt"
	"github.com/spf13/cast"
	"math"
	"math/rand"
	"sort"
	"time"
)


type Player struct{

	bet_money int
	bet_number int
	income int

}

func main(){

/*
	sum := 0.0
	p_rich := 0.0043478
	n_round := 23
	for i:=0;i<n_round;i++{
		tmp := cal(n_round,i)
		sum += tmp *math.Pow(p_rich,cast.ToFloat64(i))*math.Pow(1-p_rich,cast.ToFloat64(n_round-i))
		fmt.Println(tmp *math.Pow(p_rich,cast.ToFloat64(i))*math.Pow(1-p_rich,cast.ToFloat64(n_round-i)))
	}
	fmt.Println(sum)
*/
	//sum := 0.
	/*
	p_rich := 0.04
	n_round := 25
	tmp := cast.ToFloat64(cal(n_round,1))
	fmt.Println(tmp *math.Pow(p_rich,cast.ToFloat64(1))*math.Pow(1-p_rich,cast.ToFloat64(n_round-1)))
	*/
	//fmt.Println(probability_arr)
	//fmt.Println(probability_total)
	//random_num := randomNumGenerator()


	player_total := 30
	player_rich := 1
	player_poor := player_total - player_rich
	bet_rich := 5000
	bet_poor := 500
	banker_income := 0

	rand.Seed(time.Now().UnixNano())

	bet_matrix := [34][]Player{}
	bet_arr := [34]int{}

	for i:=0;i<player_rich;i++{
		random_num := rand.Intn(34)
		//player_list = append(player_list,)
		bet_matrix[random_num] =append(bet_matrix[random_num],Player{bet_rich,random_num,0})
	}

	for i:=0;i<player_poor;i++{
		random_num := rand.Intn(34)
		//player_list = append(player_list,Player{bet_poor,random_num ,0})
		bet_matrix[random_num] =append(bet_matrix[random_num],Player{bet_poor,random_num,0})
	}

	for i:= 0;i<34;i++{
		sum := 0
		for _,val:= range bet_matrix[i]{
			sum += val.bet_money
		}
		bet_arr[i] = sum
	}

	type kv struct {
		Key   string
		Value int
	}

	var bet_map []kv
	for k, v := range bet_arr{
		bet_map= append(bet_map, kv{cast.ToString(k), v})
	}

	sort.Slice(bet_map, func(i, j int) bool {
		return bet_map[i].Value > bet_map[j].Value  // 降序
		// return ss[i].Value > ss[j].Value  // 升序
	})

	for _, kv := range bet_map {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}

	probability_total:= 0
	probability_arr:= setProbability(&probability_total)
	probability_distribution := make([]string,probability_total)

	count := 0
	for i:=0;i<34;i++{
		for j:= 0;j<probability_arr[i];j++{
			probability_distribution[count] = bet_map[i].Key
			count += 1
		}
	}

	//fmt.Println(probability_distribution)
	selected_num_list := []int{}

	win_big := 0

	for round :=0;round<100;round++{
		selected_num := cast.ToInt(probability_distribution[rand.Intn(probability_total)])

		selected_num_list = append(selected_num_list,selected_num)

		fmt.Printf("round_numer = %v ",round)
		fmt.Println("selected number = ",selected_num)

		for i:=0;i<34;i++{
			for j:=0;j<len(bet_matrix[i]);j++{
				bet_matrix[i][j].income -= bet_matrix[i][j].bet_money
				banker_income += bet_matrix[i][j].bet_money
			}
		}

		for i:=0;i<len(bet_matrix[selected_num]);i++{
			bet_matrix[selected_num][i].income += 28*bet_matrix[selected_num][i].bet_money
			banker_income -= 28*bet_matrix[selected_num][i].bet_money

			if bet_matrix[selected_num][i].bet_money==bet_rich{
				win_big += 1
			}

		}

		fmt.Println(bet_matrix)
		fmt.Println(banker_income)
		fmt.Println(win_big)
	}

	/*
	for _, kv := range bet_map {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}
	sort.Ints(selected_num_list)
	fmt.Println(selected_num_list)*/

}



func randomNumGenerator()[]int{

	random_num := []int{}
	for i:=0;i<34;i++{
		random_num = append(random_num,i)
	}

	rand.Seed(time.Now().UnixNano())
	for i:=0;i<34;i++{
		index := rand.Intn(34-i)+i
		tmp := random_num[i]
		random_num[i] = random_num[index]
		random_num[index] = tmp
	}
	//fmt.Println(random_num)
	return random_num

}

func setProbability(probability_total *int)[]int{
	/*
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	*/

	const num = 35
	dev_container := [num]float64{}
	probability_arr := []int{}
	for i:= 0;i<num;i++{
		dev_container[i] = -2.0 + 4.0/(num-1.0)* cast.ToFloat64(i)
	}


	mean := 0.0
	dev := 1.0

	for i:=1;i<len(dev_container);i++{

		interval := 0.001
		X := dev_container[i-1]+interval*(dev_container[i] - dev_container[i-1])
		sum := 0.0
		last_p := norDistribution(mean,dev,dev_container[i-1])
		now_p:= 0.0

		for j:=0;j<1000;j++{
			now_p = norDistribution(mean,dev,X)
			sum +=(now_p +last_p )/2 *interval*(dev_container[i]- dev_container[i-1])
			X += interval*(dev_container[i] - dev_container[i-1])
			last_p = now_p

		}

		//fmt.Println(i)
		//fmt.Printf("dev=%v,sum= %v \n",dev_container[i],sum*1000.0)

		int_sum := cast.ToInt(math.Round(sum*1000))
		*probability_total += int_sum
		probability_arr = append(probability_arr,int_sum)
		//file.WriteString(cast.ToString(sum*100.0)+"\n")
	}
	sort.Sort(sort.Reverse(sort.IntSlice(probability_arr)))
	//fmt.Println(total)
	return probability_arr

}


func norDistribution(mean,dev,X float64)float64{

	p := 1/(math.Sqrt(2*math.Pi)*dev)*math.Exp(-0.5*math.Pow((X-mean)/dev,2))
	return p

}








/*
func cal(n,x int)float64{

	up := 1
	down := 1
	round := x

	for i:=0;i<round;i++{
		up *= n
		n -= 1
		down *= x
		x -= 1
	}
	return cast.ToFloat64(up) / cast.ToFloat64(down)
}*/