import React, { Component } from 'react'
import MatterService from '../services/MatterService'

class ViewMatterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            matter: {}
        }
    }

    componentDidMount(){
        MatterService.getMatterById(this.state.id).then( res => {
            this.setState({matter: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Matter Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> matterName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.matter.matterName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> leadCounsel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.matter.leadCounsel }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> MatterType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.matter.matterType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.matter.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMatterComponent
