import React, { Component } from 'react'
import SalaryComponentService from '../services/SalaryComponentService'

class ViewSalaryComponentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            salaryComponent: {}
        }
    }

    componentDidMount(){
        SalaryComponentService.getSalaryComponentById(this.state.id).then( res => {
            this.setState({salaryComponent: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View SalaryComponent Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salaryComponent.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> recurring:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salaryComponent.recurring }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ComponentType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salaryComponent.componentType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSalaryComponentComponent
