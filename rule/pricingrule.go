package rule

type Rulable interface{
	Evaluate()error
}

type PricingRule struct{

}





type PricingRules []PricingRule