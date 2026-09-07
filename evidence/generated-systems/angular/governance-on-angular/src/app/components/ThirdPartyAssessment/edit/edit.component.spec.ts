
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditThirdPartyAssessmentComponent } from './edit.component';
import { ThirdPartyAssessmentService } from '../../../services/ThirdPartyAssessment.service';

describe('EditThirdPartyAssessmentComponent', () => {
  let component: EditThirdPartyAssessmentComponent;
  let fixture: ComponentFixture<EditThirdPartyAssessmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditThirdPartyAssessmentComponent
      ],
      providers: [
        ThirdPartyAssessmentService,
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

    fixture = TestBed.createComponent(EditThirdPartyAssessmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});