
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditExpirationPolicyComponent } from './edit.component';
import { ExpirationPolicyService } from '../../../services/ExpirationPolicy.service';

describe('EditExpirationPolicyComponent', () => {
  let component: EditExpirationPolicyComponent;
  let fixture: ComponentFixture<EditExpirationPolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditExpirationPolicyComponent
      ],
      providers: [
        ExpirationPolicyService,
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

    fixture = TestBed.createComponent(EditExpirationPolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});