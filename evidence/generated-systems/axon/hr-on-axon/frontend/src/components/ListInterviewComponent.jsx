import React, { Component } from 'react'
import InterviewService from '../services/InterviewService'

class ListInterviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                interviews: []
        }
        this.addInterview = this.addInterview.bind(this);
        this.editInterview = this.editInterview.bind(this);
        this.deleteInterview = this.deleteInterview.bind(this);
    }

    deleteInterview(id){
        InterviewService.deleteInterview(id).then( res => {
            this.setState({interviews: this.state.interviews.filter(interview => interview.interviewId !== id)});
        });
    }
    viewInterview(id){
        this.props.history.push(`/view-interview/${id}`);
    }
    editInterview(id){
        this.props.history.push(`/add-interview/${id}`);
    }

    componentDidMount(){
        InterviewService.getInterviews().then((res) => {
            this.setState({ interviews: res.data});
        });
    }

    addInterview(){
        this.props.history.push('/add-interview/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Interview List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInterview}> Add Interview</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> InterviewDate </th>
                                    <th> Feedback </th>
                                    <th> Stage </th>
                                    <th> Result </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.interviews.map(
                                        interview => 
                                        <tr key = {interview.interviewId}>
                                             <td> { interview.interviewDate } </td>
                                             <td> { interview.feedback } </td>
                                             <td> { interview.stage } </td>
                                             <td> { interview.result } </td>
                                             <td>
                                                 <button onClick={ () => this.editInterview(interview.interviewId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInterview(interview.interviewId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInterview(interview.interviewId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListInterviewComponent
