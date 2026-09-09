import React, { Component } from 'react'
import ProductionLineService from '../services/ProductionLineService'

class ViewProductionLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            productionLine: {}
        }
    }

    componentDidMount(){
        ProductionLineService.getProductionLineById(this.state.id).then( res => {
            this.setState({productionLine: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ProductionLine Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productionLine.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lineCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productionLine.lineCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> LineType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productionLine.lineType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewProductionLineComponent
