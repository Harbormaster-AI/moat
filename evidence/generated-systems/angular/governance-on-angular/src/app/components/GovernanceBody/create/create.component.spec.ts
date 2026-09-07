
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateGovernanceBodyComponent } from './create.component';
import { GovernanceBodyService } from '../../../services/GovernanceBody.service';
import { Router } from '@angular/router';

describe('CreateGovernanceBodyComponent', () => {
  let component: CreateGovernanceBodyComponent;
  let fixture: ComponentFixture<CreateGovernanceBodyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateGovernanceBodyComponent
      ],
      providers: [
        GovernanceBodyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateGovernanceBodyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});