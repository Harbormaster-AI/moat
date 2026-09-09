import React, { Component } from 'react'
import InsurerService from '../services/InsurerService'

class ViewInsurerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            insurer: {}
        }
    }

    componentDidMount(){
        InsurerService.getInsurerById(this.state.id).then( res => {
            this.setState({insurer: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Insurer Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insurer.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> legalName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insurer.legalName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> domicileCountry:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insurer.domicileCountry }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> naicNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insurer.naicNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> website:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insurer.website }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInsurerComponent
