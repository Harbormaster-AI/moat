
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateJobRequisitionComponent } from './create.component';
import { JobRequisitionService } from '../../../services/JobRequisition.service';
import { Router } from '@angular/router';

describe('CreateJobRequisitionComponent', () => {
  let component: CreateJobRequisitionComponent;
  let fixture: ComponentFixture<CreateJobRequisitionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateJobRequisitionComponent
      ],
      providers: [
        JobRequisitionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateJobRequisitionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});