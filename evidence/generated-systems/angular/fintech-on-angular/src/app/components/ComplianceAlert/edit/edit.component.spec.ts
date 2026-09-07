
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditComplianceAlertComponent } from './edit.component';
import { ComplianceAlertService } from '../../../services/ComplianceAlert.service';

describe('EditComplianceAlertComponent', () => {
  let component: EditComplianceAlertComponent;
  let fixture: ComponentFixture<EditComplianceAlertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditComplianceAlertComponent
      ],
      providers: [
        ComplianceAlertService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditComplianceAlertComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});