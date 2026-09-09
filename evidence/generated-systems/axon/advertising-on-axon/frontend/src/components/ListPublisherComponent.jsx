import React, { Component } from 'react'
import PublisherService from '../services/PublisherService'

class ListPublisherComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                publishers: []
        }
        this.addPublisher = this.addPublisher.bind(this);
        this.editPublisher = this.editPublisher.bind(this);
        this.deletePublisher = this.deletePublisher.bind(this);
    }

    deletePublisher(id){
        PublisherService.deletePublisher(id).then( res => {
            this.setState({publishers: this.state.publishers.filter(publisher => publisher.publisherId !== id)});
        });
    }
    viewPublisher(id){
        this.props.history.push(`/view-publisher/${id}`);
    }
    editPublisher(id){
        this.props.history.push(`/add-publisher/${id}`);
    }

    componentDidMount(){
        PublisherService.getPublishers().then((res) => {
            this.setState({ publishers: res.data});
        });
    }

    addPublisher(){
        this.props.history.push('/add-publisher/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Publisher List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPublisher}> Add Publisher</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Website </th>
                                    <th> PublisherType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.publishers.map(
                                        publisher => 
                                        <tr key = {publisher.publisherId}>
                                             <td> { publisher.name } </td>
                                             <td> { publisher.website } </td>
                                             <td> { publisher.publisherType } </td>
                                             <td>
                                                 <button onClick={ () => this.editPublisher(publisher.publisherId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePublisher(publisher.publisherId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPublisher(publisher.publisherId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPublisherComponent
