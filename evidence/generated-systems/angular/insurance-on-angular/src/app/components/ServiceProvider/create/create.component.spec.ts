
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateServiceProviderComponent } from './create.component';
import { ServiceProviderService } from '../../../services/ServiceProvider.service';
import { Router } from '@angular/router';

describe('CreateServiceProviderComponent', () => {
  let component: CreateServiceProviderComponent;
  let fixture: ComponentFixture<CreateServiceProviderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateServiceProviderComponent
      ],
      providers: [
        ServiceProviderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateServiceProviderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});