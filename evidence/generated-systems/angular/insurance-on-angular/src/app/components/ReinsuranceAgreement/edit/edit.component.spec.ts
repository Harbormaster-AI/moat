
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditReinsuranceAgreementComponent } from './edit.component';
import { ReinsuranceAgreementService } from '../../../services/ReinsuranceAgreement.service';

describe('EditReinsuranceAgreementComponent', () => {
  let component: EditReinsuranceAgreementComponent;
  let fixture: ComponentFixture<EditReinsuranceAgreementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditReinsuranceAgreementComponent
      ],
      providers: [
        ReinsuranceAgreementService,
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

    fixture = TestBed.createComponent(EditReinsuranceAgreementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});