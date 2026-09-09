import React, { Component } from 'react'
import PublisherService from '../services/PublisherService'

class ViewPublisherComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            publisher: {}
        }
    }

    componentDidMount(){
        PublisherService.getPublisherById(this.state.id).then( res => {
            this.setState({publisher: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Publisher Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.publisher.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> website:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.publisher.website }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PublisherType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.publisher.publisherType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPublisherComponent
