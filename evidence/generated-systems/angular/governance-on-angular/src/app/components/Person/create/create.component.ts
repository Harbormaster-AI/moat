import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PersonService } from '../../../services/Person.service';
import { Person } from '../../../models/Person';
import { SubBaseComponent } from '../../Person/sub.base.component';

@Component({
    selector: 'app-create-person',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePersonComponent extends SubBaseComponent implements OnInit {

    title = 'Add Person';

    personForm: FormGroup;
    person: Person;

    constructor( http: HttpClient,
        private personService: PersonService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPerson(firstName, lastName, email, department, RoleAssignments, OwnedPolicies, CorrectiveActions): void {
        this.personService
        .addPerson(firstName, lastName, email, department, RoleAssignments, OwnedPolicies, CorrectiveActions)
            .subscribe(() => {
                this.router.navigate(['/indexPerson']);
            });
    }

    ngOnInit(): void {
    }
}