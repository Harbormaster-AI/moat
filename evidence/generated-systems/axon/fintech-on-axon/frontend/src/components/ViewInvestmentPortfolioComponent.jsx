import React, { Component } from 'react'
import InvestmentPortfolioService from '../services/InvestmentPortfolioService'

class ViewInvestmentPortfolioComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            investmentPortfolio: {}
        }
    }

    componentDidMount(){
        InvestmentPortfolioService.getInvestmentPortfolioById(this.state.id).then( res => {
            this.setState({investmentPortfolio: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InvestmentPortfolio Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> portfolioCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.investmentPortfolio.portfolioCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> baseCurrency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.investmentPortfolio.baseCurrency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> createdAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.investmentPortfolio.createdAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.investmentPortfolio.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInvestmentPortfolioComponent
