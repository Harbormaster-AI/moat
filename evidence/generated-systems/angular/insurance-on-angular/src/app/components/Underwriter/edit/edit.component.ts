import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { UnderwriterService } from '../../../services/Underwriter.service';
import { SubBaseComponent } from '../../Underwriter/sub.base.component';


@Component({
    selector: 'app-edit-underwriter',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditUnderwriterComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Underwriter';

    underwriterForm: FormGroup;
    underwriter: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: UnderwriterService,
        private fb: FormBuilder
) {
        super(http);
        this.underwriterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      employeeId: ['', Validators.required],
      authorityLimit: ['', Validators.required],
      Decisions: ['', ],
      Insurer: ['', ]
        });
    }

    
    updateUnderwriter(firstName, lastName, employeeId, authorityLimit, Decisions, Insurer): void {
        this.route.params.subscribe((params) => {

                        this.service.updateUnderwriter(firstName, lastName, employeeId, authorityLimit, Decisions, Insurer, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexUnderwriter']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getUnderwriter(params['id']).subscribe(res => {
                this.underwriter = res;
            });
        });
    }
}