import React, { Component } from 'react'
import CabinLayoutService from '../services/CabinLayoutService'

class ViewCabinLayoutComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            cabinLayout: {}
        }
    }

    componentDidMount(){
        CabinLayoutService.getCabinLayoutById(this.state.id).then( res => {
            this.setState({cabinLayout: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CabinLayout Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> layoutCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cabinLayout.layoutCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalSeats:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cabinLayout.totalSeats }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> classLayout:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cabinLayout.classLayout }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCabinLayoutComponent
