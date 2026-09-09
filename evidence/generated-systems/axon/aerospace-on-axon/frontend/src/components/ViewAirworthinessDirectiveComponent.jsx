import React, { Component } from 'react'
import AirworthinessDirectiveService from '../services/AirworthinessDirectiveService'

class ViewAirworthinessDirectiveComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            airworthinessDirective: {}
        }
    }

    componentDidMount(){
        AirworthinessDirectiveService.getAirworthinessDirectiveById(this.state.id).then( res => {
            this.setState({airworthinessDirective: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AirworthinessDirective Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> directiveNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.airworthinessDirective.directiveNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.airworthinessDirective.title }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAirworthinessDirectiveComponent
