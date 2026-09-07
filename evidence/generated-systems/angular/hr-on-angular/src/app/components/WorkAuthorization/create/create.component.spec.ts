
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateWorkAuthorizationComponent } from './create.component';
import { WorkAuthorizationService } from '../../../services/WorkAuthorization.service';
import { Router } from '@angular/router';

describe('CreateWorkAuthorizationComponent', () => {
  let component: CreateWorkAuthorizationComponent;
  let fixture: ComponentFixture<CreateWorkAuthorizationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateWorkAuthorizationComponent
      ],
      providers: [
        WorkAuthorizationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateWorkAuthorizationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});