import React, { Component } from 'react'
import SubscriberService from '../services/SubscriberService'

class ListSubscriberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                subscribers: []
        }
        this.addSubscriber = this.addSubscriber.bind(this);
        this.editSubscriber = this.editSubscriber.bind(this);
        this.deleteSubscriber = this.deleteSubscriber.bind(this);
    }

    deleteSubscriber(id){
        SubscriberService.deleteSubscriber(id).then( res => {
            this.setState({subscribers: this.state.subscribers.filter(subscriber => subscriber.subscriberId !== id)});
        });
    }
    viewSubscriber(id){
        this.props.history.push(`/view-subscriber/${id}`);
    }
    editSubscriber(id){
        this.props.history.push(`/add-subscriber/${id}`);
    }

    componentDidMount(){
        SubscriberService.getSubscribers().then((res) => {
            this.setState({ subscribers: res.data});
        });
    }

    addSubscriber(){
        this.props.history.push('/add-subscriber/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Subscriber List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSubscriber}> Add Subscriber</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Address </th>
                                    <th> Channel </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.subscribers.map(
                                        subscriber => 
                                        <tr key = {subscriber.subscriberId}>
                                             <td> { subscriber.name } </td>
                                             <td> { subscriber.address } </td>
                                             <td> { subscriber.channel } </td>
                                             <td>
                                                 <button onClick={ () => this.editSubscriber(subscriber.subscriberId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSubscriber(subscriber.subscriberId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSubscriber(subscriber.subscriberId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSubscriberComponent
