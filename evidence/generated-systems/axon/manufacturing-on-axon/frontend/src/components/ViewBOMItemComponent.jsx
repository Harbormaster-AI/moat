import React, { Component } from 'react'
import BOMItemService from '../services/BOMItemService'

class ViewBOMItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            bOMItem: {}
        }
    }

    componentDidMount(){
        BOMItemService.getBOMItemById(this.state.id).then( res => {
            this.setState({bOMItem: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BOMItem Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lineNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bOMItem.lineNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bOMItem.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> scrapPercent:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bOMItem.scrapPercent }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBOMItemComponent
