
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditPolicyAcknowledgementComponent } from './edit.component';
import { PolicyAcknowledgementService } from '../../../services/PolicyAcknowledgement.service';

describe('EditPolicyAcknowledgementComponent', () => {
  let component: EditPolicyAcknowledgementComponent;
  let fixture: ComponentFixture<EditPolicyAcknowledgementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditPolicyAcknowledgementComponent
      ],
      providers: [
        PolicyAcknowledgementService,
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

    fixture = TestBed.createComponent(EditPolicyAcknowledgementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});