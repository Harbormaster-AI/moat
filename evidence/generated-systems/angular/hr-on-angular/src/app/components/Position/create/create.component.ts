import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PositionService } from '../../../services/Position.service';
import { Position } from '../../../models/Position';
import { SubBaseComponent } from '../../Position/sub.base.component';

@Component({
    selector: 'app-create-position',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePositionComponent extends SubBaseComponent implements OnInit {

    title = 'Add Position';

    positionForm: FormGroup;
    position: Position;

    constructor( http: HttpClient,
        private positionService: PositionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.positionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  positionCode: ['', Validators.required],
      fte: ['', Validators.required],
      Department: ['', ],
      JobProfile: ['', ],
      CostCenter: ['', ],
      Location: ['', ],
      ManagerPosition: ['', ],
      DirectReports: ['', ],
      Assignments: ['', ],
      Status: ['', ],
      WorkLocationType: ['', ]
        });
    }

    
    addPosition(positionCode, fte, Department, JobProfile, CostCenter, Location, ManagerPosition, DirectReports, Assignments, Status, WorkLocationType): void {
        this.positionService
        .addPosition(positionCode, fte, Department, JobProfile, CostCenter, Location, ManagerPosition, DirectReports, Assignments, Status, WorkLocationType)
            .subscribe(() => {
                this.router.navigate(['/indexPosition']);
            });
    }

    ngOnInit(): void {
    }
}