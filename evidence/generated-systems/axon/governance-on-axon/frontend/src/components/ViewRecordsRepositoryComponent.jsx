import React, { Component } from 'react'
import RecordsRepositoryService from '../services/RecordsRepositoryService'

class ViewRecordsRepositoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            recordsRepository: {}
        }
    }

    componentDidMount(){
        RecordsRepositoryService.getRecordsRepositoryById(this.state.id).then( res => {
            this.setState({recordsRepository: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View RecordsRepository Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.recordsRepository.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> location:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.recordsRepository.location }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ownerDepartment:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.recordsRepository.ownerDepartment }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> RepositoryType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.recordsRepository.repositoryType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRecordsRepositoryComponent
