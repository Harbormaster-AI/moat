import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PersonService } from '../../../services/Person.service';
import { SubBaseComponent } from '../../Person/sub.base.component';


@Component({
    selector: 'app-edit-person',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPersonComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Person';

    personForm: FormGroup;
    person: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PersonService,
        private fb: FormBuilder
) {
        super(http);
        this.personForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      email: ['', Validators.required],
      department: ['', Validators.required],
      RoleAssignments: ['', ],
      OwnedPolicies: ['', ],
      CorrectiveActions: ['', ]
        });
    }

    
    updatePerson(firstName, lastName, email, department, RoleAssignments, OwnedPolicies, CorrectiveActions): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePerson(firstName, lastName, email, department, RoleAssignments, OwnedPolicies, CorrectiveActions, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPerson']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPerson(params['id']).subscribe(res => {
                this.person = res;
            });
        });
    }
}