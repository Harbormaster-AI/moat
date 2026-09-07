import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { APIClientService } from '../../../services/APIClient.service';
import { SubBaseComponent } from '../../APIClient/sub.base.component';


@Component({
    selector: 'app-edit-aPIClient',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAPIClientComponent extends SubBaseComponent implements OnInit {

    title = 'Edit APIClient';

    aPIClientForm: FormGroup;
    aPIClient: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: APIClientService,
        private fb: FormBuilder
) {
        super(http);
        this.aPIClientForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      clientId: ['', Validators.required],
      redirectUri: ['', Validators.required],
      Consents: ['', ],
      ClientType: ['', ]
        });
    }

    
    updateAPIClient(name, clientId, redirectUri, Consents, ClientType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAPIClient(name, clientId, redirectUri, Consents, ClientType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAPIClient']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAPIClient(params['id']).subscribe(res => {
                this.aPIClient = res;
            });
        });
    }
}