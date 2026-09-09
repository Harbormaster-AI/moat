import React, { Component } from 'react'
import DependentService from '../services/DependentService'

class ViewDependentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dependent: {}
        }
    }

    componentDidMount(){
        DependentService.getDependentById(this.state.id).then( res => {
            this.setState({dependent: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Dependent Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> firstName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dependent.firstName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lastName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dependent.lastName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> birthDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dependent.birthDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Relationship:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dependent.relationship }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDependentComponent
