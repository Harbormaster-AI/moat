
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditGovernanceBodyComponent } from './edit.component';
import { GovernanceBodyService } from '../../../services/GovernanceBody.service';

describe('EditGovernanceBodyComponent', () => {
  let component: EditGovernanceBodyComponent;
  let fixture: ComponentFixture<EditGovernanceBodyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditGovernanceBodyComponent
      ],
      providers: [
        GovernanceBodyService,
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

    fixture = TestBed.createComponent(EditGovernanceBodyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});