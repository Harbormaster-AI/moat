import React, { Component } from 'react'
import BackgroundCheckService from '../services/BackgroundCheckService'

class ListBackgroundCheckComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                backgroundChecks: []
        }
        this.addBackgroundCheck = this.addBackgroundCheck.bind(this);
        this.editBackgroundCheck = this.editBackgroundCheck.bind(this);
        this.deleteBackgroundCheck = this.deleteBackgroundCheck.bind(this);
    }

    deleteBackgroundCheck(id){
        BackgroundCheckService.deleteBackgroundCheck(id).then( res => {
            this.setState({backgroundChecks: this.state.backgroundChecks.filter(backgroundCheck => backgroundCheck.backgroundCheckId !== id)});
        });
    }
    viewBackgroundCheck(id){
        this.props.history.push(`/view-backgroundCheck/${id}`);
    }
    editBackgroundCheck(id){
        this.props.history.push(`/add-backgroundCheck/${id}`);
    }

    componentDidMount(){
        BackgroundCheckService.getBackgroundChecks().then((res) => {
            this.setState({ backgroundChecks: res.data});
        });
    }

    addBackgroundCheck(){
        this.props.history.push('/add-backgroundCheck/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BackgroundCheck List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBackgroundCheck}> Add BackgroundCheck</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CheckNumber </th>
                                    <th> Provider </th>
                                    <th> CompletedDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.backgroundChecks.map(
                                        backgroundCheck => 
                                        <tr key = {backgroundCheck.backgroundCheckId}>
                                             <td> { backgroundCheck.checkNumber } </td>
                                             <td> { backgroundCheck.provider } </td>
                                             <td> { backgroundCheck.completedDate } </td>
                                             <td> { backgroundCheck.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editBackgroundCheck(backgroundCheck.backgroundCheckId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBackgroundCheck(backgroundCheck.backgroundCheckId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBackgroundCheck(backgroundCheck.backgroundCheckId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBackgroundCheckComponent
