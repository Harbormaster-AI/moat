import React, { Component } from 'react'
import PublisherService from '../services/PublisherService';

class UpdatePublisherComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                website: '',
                publisherType: ''
        }
        this.updatePublisher = this.updatePublisher.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changePublisherTypeHandler = this.changePublisherTypeHandler.bind(this);
    }

    componentDidMount(){
        PublisherService.getPublisherById(this.state.id).then( (res) =>{
            let publisher = res.data;
            this.setState({
                name: publisher.name,
                website: publisher.website,
                publisherType: publisher.publisherType
            });
        });
    }

    updatePublisher = (e) => {
        e.preventDefault();
        let publisher = {
            publisherId: this.state.id,
            name: this.state.name,
            website: this.state.website,
            publisherType: this.state.publisherType
        };
        console.log('publisher => ' + JSON.stringify(publisher));
        console.log('id => ' + JSON.stringify(this.state.id));
        PublisherService.updatePublisher(publisher).then( res => {
            this.props.history.push('/publishers');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }
    changePublisherTypeHandler= (event) => {
        this.setState({publisherType: event.target.value});
    }

    cancel(){
        this.props.history.push('/publishers');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Publisher</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> PublisherType: </label>
                                                <select value={this.state.publisherType} onChange={this.changePublisherTypeHandler}>
                      <option name="PublisherType" className="form-control" >
                          Site
                      </option>
                      <option name="PublisherType" className="form-control" >
                          App
                      </option>
                      <option name="PublisherType" className="form-control" >
                          Network
                      </option>
                      <option name="PublisherType" className="form-control" >
                          CTVApp
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePublisher}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdatePublisherComponent
