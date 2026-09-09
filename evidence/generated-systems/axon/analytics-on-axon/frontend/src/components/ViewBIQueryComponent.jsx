import React, { Component } from 'react'
import BIQueryService from '../services/BIQueryService'

class ViewBIQueryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            bIQuery: {}
        }
    }

    componentDidMount(){
        BIQueryService.getBIQueryById(this.state.id).then( res => {
            this.setState({bIQuery: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BIQuery Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bIQuery.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> text:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bIQuery.text }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Dialect:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bIQuery.dialect }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBIQueryComponent
