package client


func ExecuteArb(neg_risk_id string){
	markets_to_trade := NegRiskMarketMap[neg_risk_id]
	volume_to_trade := 10e10
	for _, market := range markets_to_trade {
		
}