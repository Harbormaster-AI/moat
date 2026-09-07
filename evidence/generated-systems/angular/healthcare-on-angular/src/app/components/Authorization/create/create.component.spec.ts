
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAuthorizationComponent } from './create.component';
import { AuthorizationService } from '../../../services/Authorization.service';
import { Router } from '@angular/router';

describe('CreateAuthorizationComponent', () => {
  let component: CreateAuthorizationComponent;
  let fixture: ComponentFixture<CreateAuthorizationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAuthorizationComponent
      ],
      providers: [
        AuthorizationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAuthorizationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});