import React, { Component } from 'react'
import ScreeningService from '../services/ScreeningService'

class ViewScreeningComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            screening: {}
        }
    }

    componentDidMount(){
        ScreeningService.getScreeningById(this.state.id).then( res => {
            this.setState({screening: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Screening Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> score:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.screening.score }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> screenedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.screening.screenedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ScreeningType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.screening.screeningType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.screening.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewScreeningComponent
