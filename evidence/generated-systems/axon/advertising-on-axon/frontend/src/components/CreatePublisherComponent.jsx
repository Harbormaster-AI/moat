import React, { Component } from 'react'
import PublisherService from '../services/PublisherService';

class CreatePublisherComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                website: '',
                publisherType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changePublisherTypeHandler = this.changePublisherTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PublisherService.getPublisherById(this.state.id).then( (res) =>{
                let publisher = res.data;
                this.setState({
                    name: publisher.name,
                    website: publisher.website,
                    publisherType: publisher.publisherType
                });
            });
        }        
    }
    saveOrUpdatePublisher = (e) => {
        e.preventDefault();
        let publisher = {
                publisherId: this.state.id,
                name: this.state.name,
                website: this.state.website,
                publisherType: this.state.publisherType
            };
        console.log('publisher => ' + JSON.stringify(publisher));

        // step 5
        if(this.state.id === '_add'){
            publisher.publisherId=''
            PublisherService.createPublisher(publisher).then(res =>{
                this.props.history.push('/publishers');
            });
        }else{
            PublisherService.updatePublisher(publisher).then( res => {
                this.props.history.push('/publishers');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Publisher</h3>
        }else{
            return <h3 className="text-center">Update Publisher</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> PublisherType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePublisher}>Save</button>
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

export default CreatePublisherComponent
