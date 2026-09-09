import React, { Component } from 'react'
import StockKeepingUnitService from '../services/StockKeepingUnitService'

class ViewStockKeepingUnitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            stockKeepingUnit: {}
        }
    }

    componentDidMount(){
        StockKeepingUnitService.getStockKeepingUnitById(this.state.id).then( res => {
            this.setState({stockKeepingUnit: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View StockKeepingUnit Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> skuCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockKeepingUnit.skuCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockKeepingUnit.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> weight:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockKeepingUnit.weight }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> weightUnit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockKeepingUnit.weightUnit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> volume:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockKeepingUnit.volume }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> volumeUnit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockKeepingUnit.volumeUnit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shelfLifeDays:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockKeepingUnit.shelfLifeDays }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> hazardousMaterial:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockKeepingUnit.hazardousMaterial }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ItemType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockKeepingUnit.itemType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> UnitOfMeasure:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockKeepingUnit.unitOfMeasure }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewStockKeepingUnitComponent
