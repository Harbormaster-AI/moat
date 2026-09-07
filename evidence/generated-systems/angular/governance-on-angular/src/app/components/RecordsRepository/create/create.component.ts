import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RecordsRepositoryService } from '../../../services/RecordsRepository.service';
import { RecordsRepository } from '../../../models/RecordsRepository';
import { SubBaseComponent } from '../../RecordsRepository/sub.base.component';

@Component({
    selector: 'app-create-recordsRepository',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRecordsRepositoryComponent extends SubBaseComponent implements OnInit {

    title = 'Add RecordsRepository';

    recordsRepositoryForm: FormGroup;
    recordsRepository: RecordsRepository;

    constructor( http: HttpClient,
        private recordsRepositoryService: RecordsRepositoryService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.recordsRepositoryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      location: ['', Validators.required],
      ownerDepartment: ['', Validators.required],
      Organization: ['', ],
      Records: ['', ],
      Systems: ['', ],
      RetentionSchedules: ['', ],
      LegalHolds: ['', ],
      RepositoryType: ['', ]
        });
    }

    
    addRecordsRepository(name, location, ownerDepartment, Organization, Records, Systems, RetentionSchedules, LegalHolds, RepositoryType): void {
        this.recordsRepositoryService
        .addRecordsRepository(name, location, ownerDepartment, Organization, Records, Systems, RetentionSchedules, LegalHolds, RepositoryType)
            .subscribe(() => {
                this.router.navigate(['/indexRecordsRepository']);
            });
    }

    ngOnInit(): void {
    }
}