
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateThirdPartyAssessmentComponent } from './create.component';
import { ThirdPartyAssessmentService } from '../../../services/ThirdPartyAssessment.service';
import { Router } from '@angular/router';

describe('CreateThirdPartyAssessmentComponent', () => {
  let component: CreateThirdPartyAssessmentComponent;
  let fixture: ComponentFixture<CreateThirdPartyAssessmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateThirdPartyAssessmentComponent
      ],
      providers: [
        ThirdPartyAssessmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateThirdPartyAssessmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});