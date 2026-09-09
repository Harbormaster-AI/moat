import React, { Component } from 'react'
import NoteService from '../services/NoteService'

class ViewNoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            note: {}
        }
    }

    componentDidMount(){
        NoteService.getNoteById(this.state.id).then( res => {
            this.setState({note: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Note Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.note.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> content:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.note.content }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> createdAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.note.createdAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> updatedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.note.updatedAt }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewNoteComponent
