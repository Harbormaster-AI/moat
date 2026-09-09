import React, { Component } from 'react'
import JobFamilyService from '../services/JobFamilyService'

class ViewJobFamilyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            jobFamily: {}
        }
    }

    componentDidMount(){
        JobFamilyService.getJobFamilyById(this.state.id).then( res => {
            this.setState({jobFamily: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View JobFamily Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobFamily.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobFamily.description }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewJobFamilyComponent
