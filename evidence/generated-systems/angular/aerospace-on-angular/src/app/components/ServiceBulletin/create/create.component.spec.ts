
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateServiceBulletinComponent } from './create.component';
import { ServiceBulletinService } from '../../../services/ServiceBulletin.service';
import { Router } from '@angular/router';

describe('CreateServiceBulletinComponent', () => {
  let component: CreateServiceBulletinComponent;
  let fixture: ComponentFixture<CreateServiceBulletinComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateServiceBulletinComponent
      ],
      providers: [
        ServiceBulletinService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateServiceBulletinComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});