import React, { Component } from 'react'
import UsageLimitService from '../services/UsageLimitService'

class ListUsageLimitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                usageLimits: []
        }
        this.addUsageLimit = this.addUsageLimit.bind(this);
        this.editUsageLimit = this.editUsageLimit.bind(this);
        this.deleteUsageLimit = this.deleteUsageLimit.bind(this);
    }

    deleteUsageLimit(id){
        UsageLimitService.deleteUsageLimit(id).then( res => {
            this.setState({usageLimits: this.state.usageLimits.filter(usageLimit => usageLimit.usageLimitId !== id)});
        });
    }
    viewUsageLimit(id){
        this.props.history.push(`/view-usageLimit/${id}`);
    }
    editUsageLimit(id){
        this.props.history.push(`/add-usageLimit/${id}`);
    }

    componentDidMount(){
        UsageLimitService.getUsageLimits().then((res) => {
            this.setState({ usageLimits: res.data});
        });
    }

    addUsageLimit(){
        this.props.history.push('/add-usageLimit/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">UsageLimit List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addUsageLimit}> Add UsageLimit</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Amount </th>
                                    <th> Count </th>
                                    <th> Scope </th>
                                    <th> Period </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.usageLimits.map(
                                        usageLimit => 
                                        <tr key = {usageLimit.usageLimitId}>
                                             <td> { usageLimit.name } </td>
                                             <td> { usageLimit.amount } </td>
                                             <td> { usageLimit.count } </td>
                                             <td> { usageLimit.scope } </td>
                                             <td> { usageLimit.period } </td>
                                             <td>
                                                 <button onClick={ () => this.editUsageLimit(usageLimit.usageLimitId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteUsageLimit(usageLimit.usageLimitId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewUsageLimit(usageLimit.usageLimitId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListUsageLimitComponent
