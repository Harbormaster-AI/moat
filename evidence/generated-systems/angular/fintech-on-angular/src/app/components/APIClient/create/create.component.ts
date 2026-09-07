import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { APIClientService } from '../../../services/APIClient.service';
import { APIClient } from '../../../models/APIClient';
import { SubBaseComponent } from '../../APIClient/sub.base.component';

@Component({
    selector: 'app-create-aPIClient',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAPIClientComponent extends SubBaseComponent implements OnInit {

    title = 'Add APIClient';

    aPIClientForm: FormGroup;
    aPIClient: APIClient;

    constructor( http: HttpClient,
        private aPIClientService: APIClientService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAPIClient(name, clientId, redirectUri, Consents, ClientType): void {
        this.aPIClientService
        .addAPIClient(name, clientId, redirectUri, Consents, ClientType)
            .subscribe(() => {
                this.router.navigate(['/indexAPIClient']);
            });
    }

    ngOnInit(): void {
    }
}