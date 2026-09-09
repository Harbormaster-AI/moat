import React, { Component } from 'react'
import CycleCountEntryService from '../services/CycleCountEntryService'

class ViewCycleCountEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            cycleCountEntry: {}
        }
    }

    componentDidMount(){
        CycleCountEntryService.getCycleCountEntryById(this.state.id).then( res => {
            this.setState({cycleCountEntry: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CycleCountEntry Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lineNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCountEntry.lineNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> systemQuantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCountEntry.systemQuantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> countedQuantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCountEntry.countedQuantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> varianceQuantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCountEntry.varianceQuantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> recountRequired:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCountEntry.recountRequired }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> StockStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCountEntry.stockStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCycleCountEntryComponent
