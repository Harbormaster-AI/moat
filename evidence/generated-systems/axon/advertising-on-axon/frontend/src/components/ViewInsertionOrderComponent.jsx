import React, { Component } from 'react'
import InsertionOrderService from '../services/InsertionOrderService'

class ViewInsertionOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            insertionOrder: {}
        }
    }

    componentDidMount(){
        InsertionOrderService.getInsertionOrderById(this.state.id).then( res => {
            this.setState({insertionOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InsertionOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ioNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insertionOrder.ioNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> agreedBudget:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insertionOrder.agreedBudget }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> flight:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insertionOrder.flight }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insertionOrder.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInsertionOrderComponent
