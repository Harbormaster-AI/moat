import React, { Component } from 'react'
import Exception_Service from '../services/Exception_Service'

class ViewException_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            exception_: {}
        }
    }

    componentDidMount(){
        Exception_Service.getException_ById(this.state.id).then( res => {
            this.setState({exception_: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Exception_ Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.exception_.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> justification:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.exception_.justification }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.exception_.startDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.exception_.endDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ExceptionType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.exception_.exceptionType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.exception_.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewException_Component
