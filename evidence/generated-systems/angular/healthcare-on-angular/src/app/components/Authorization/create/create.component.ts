import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthorizationService } from '../../../services/Authorization.service';
import { Authorization } from '../../../models/Authorization';
import { SubBaseComponent } from '../../Authorization/sub.base.component';

@Component({
    selector: 'app-create-authorization',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAuthorizationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Authorization';

    authorizationForm: FormGroup;
    authorization: Authorization;

    constructor( http: HttpClient,
        private authorizationService: AuthorizationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.authorizationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  authNumber: ['', Validators.required],
      requestedService: ['', Validators.required],
      Coverage: ['', ],
      Order: ['', ],
      Status: ['', ]
        });
    }

    
    addAuthorization(authNumber, requestedService, Coverage, Order, Status): void {
        this.authorizationService
        .addAuthorization(authNumber, requestedService, Coverage, Order, Status)
            .subscribe(() => {
                this.router.navigate(['/indexAuthorization']);
            });
    }

    ngOnInit(): void {
    }
}