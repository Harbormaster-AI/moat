import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UnderwriterService } from '../../../services/Underwriter.service';
import { Underwriter } from '../../../models/Underwriter';
import { SubBaseComponent } from '../../Underwriter/sub.base.component';

@Component({
    selector: 'app-create-underwriter',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateUnderwriterComponent extends SubBaseComponent implements OnInit {

    title = 'Add Underwriter';

    underwriterForm: FormGroup;
    underwriter: Underwriter;

    constructor( http: HttpClient,
        private underwriterService: UnderwriterService,
        private fb: FormBuilder,
        private router: Router
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

    
    addUnderwriter(firstName, lastName, employeeId, authorityLimit, Decisions, Insurer): void {
        this.underwriterService
        .addUnderwriter(firstName, lastName, employeeId, authorityLimit, Decisions, Insurer)
            .subscribe(() => {
                this.router.navigate(['/indexUnderwriter']);
            });
    }

    ngOnInit(): void {
    }
}