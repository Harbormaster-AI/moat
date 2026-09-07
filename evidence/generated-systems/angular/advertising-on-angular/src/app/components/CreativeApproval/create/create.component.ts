import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CreativeApprovalService } from '../../../services/CreativeApproval.service';
import { CreativeApproval } from '../../../models/CreativeApproval';
import { SubBaseComponent } from '../../CreativeApproval/sub.base.component';

@Component({
    selector: 'app-create-creativeApproval',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCreativeApprovalComponent extends SubBaseComponent implements OnInit {

    title = 'Add CreativeApproval';

    creativeApprovalForm: FormGroup;
    creativeApproval: CreativeApproval;

    constructor( http: HttpClient,
        private creativeApprovalService: CreativeApprovalService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.creativeApprovalForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  reviewer: ['', Validators.required],
      reviewedAt: ['', Validators.required],
      CreativeAsset: ['', ],
      Publisher: ['', ],
      Status: ['', ]
        });
    }

    
    addCreativeApproval(reviewer, reviewedAt, CreativeAsset, Publisher, Status): void {
        this.creativeApprovalService
        .addCreativeApproval(reviewer, reviewedAt, CreativeAsset, Publisher, Status)
            .subscribe(() => {
                this.router.navigate(['/indexCreativeApproval']);
            });
    }

    ngOnInit(): void {
    }
}