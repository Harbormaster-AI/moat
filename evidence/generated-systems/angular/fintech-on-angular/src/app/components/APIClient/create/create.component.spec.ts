
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAPIClientComponent } from './create.component';
import { APIClientService } from '../../../services/APIClient.service';
import { Router } from '@angular/router';

describe('CreateAPIClientComponent', () => {
  let component: CreateAPIClientComponent;
  let fixture: ComponentFixture<CreateAPIClientComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAPIClientComponent
      ],
      providers: [
        APIClientService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAPIClientComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});