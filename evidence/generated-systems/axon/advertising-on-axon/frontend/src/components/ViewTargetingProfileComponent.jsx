import React, { Component } from 'react'
import TargetingProfileService from '../services/TargetingProfileService'

class ViewTargetingProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            targetingProfile: {}
        }
    }

    componentDidMount(){
        TargetingProfileService.getTargetingProfileById(this.state.id).then( res => {
            this.setState({targetingProfile: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TargetingProfile Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.targetingProfile.name }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTargetingProfileComponent
