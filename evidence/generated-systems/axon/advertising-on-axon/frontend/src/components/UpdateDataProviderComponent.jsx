import React, { Component } from 'react'
import DataProviderService from '../services/DataProviderService';

class UpdateDataProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                website: '',
                providerType: ''
        }
        this.updateDataProvider = this.updateDataProvider.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changeProviderTypeHandler = this.changeProviderTypeHandler.bind(this);
    }

    componentDidMount(){
        DataProviderService.getDataProviderById(this.state.id).then( (res) =>{
            let dataProvider = res.data;
            this.setState({
                name: dataProvider.name,
                website: dataProvider.website,
                providerType: dataProvider.providerType
            });
        });
    }

    updateDataProvider = (e) => {
        e.preventDefault();
        let dataProvider = {
            dataProviderId: this.state.id,
            name: this.state.name,
            website: this.state.website,
            providerType: this.state.providerType
        };
        console.log('dataProvider => ' + JSON.stringify(dataProvider));
        console.log('id => ' + JSON.stringify(this.state.id));
        DataProviderService.updateDataProvider(dataProvider).then( res => {
            this.props.history.push('/dataProviders');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }
    changeProviderTypeHandler= (event) => {
        this.setState({providerType: event.target.value});
    }

    cancel(){
        this.props.history.push('/dataProviders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DataProvider</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> ProviderType: </label>
                                                <select value={this.state.providerType} onChange={this.changeProviderTypeHandler}>
                      <option name="ProviderType" className="form-control" >
                          FirstParty
                      </option>
                      <option name="ProviderType" className="form-control" >
                          SecondParty
                      </option>
                      <option name="ProviderType" className="form-control" >
                          ThirdParty
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDataProvider}>Save</button>
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

export default UpdateDataProviderComponent
