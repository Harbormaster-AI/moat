import React, { Component } from 'react'
import InsuredObjectService from '../services/InsuredObjectService'

class ViewInsuredObjectComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            insuredObject: {}
        }
    }

    componentDidMount(){
        InsuredObjectService.getInsuredObjectById(this.state.id).then( res => {
            this.setState({insuredObject: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InsuredObject Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insuredObject.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> serialOrId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insuredObject.serialOrId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> primaryAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insuredObject.primaryAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ObjectType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insuredObject.objectType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInsuredObjectComponent
