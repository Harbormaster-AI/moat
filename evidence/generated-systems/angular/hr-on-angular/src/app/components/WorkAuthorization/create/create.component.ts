import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { WorkAuthorizationService } from '../../../services/WorkAuthorization.service';
import { WorkAuthorization } from '../../../models/WorkAuthorization';
import { SubBaseComponent } from '../../WorkAuthorization/sub.base.component';

@Component({
    selector: 'app-create-workAuthorization',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateWorkAuthorizationComponent extends SubBaseComponent implements OnInit {

    title = 'Add WorkAuthorization';

    workAuthorizationForm: FormGroup;
    workAuthorization: WorkAuthorization;

    constructor( http: HttpClient,
        private workAuthorizationService: WorkAuthorizationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.workAuthorizationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  country: ['', Validators.required],
      expirationDate: ['', Validators.required],
      Employee: ['', ],
      Documents: ['', ],
      Status: ['', ]
        });
    }

    
    addWorkAuthorization(country, expirationDate, Employee, Documents, Status): void {
        this.workAuthorizationService
        .addWorkAuthorization(country, expirationDate, Employee, Documents, Status)
            .subscribe(() => {
                this.router.navigate(['/indexWorkAuthorization']);
            });
    }

    ngOnInit(): void {
    }
}