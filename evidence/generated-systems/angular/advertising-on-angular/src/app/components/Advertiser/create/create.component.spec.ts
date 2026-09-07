
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAdvertiserComponent } from './create.component';
import { AdvertiserService } from '../../../services/Advertiser.service';
import { Router } from '@angular/router';

describe('CreateAdvertiserComponent', () => {
  let component: CreateAdvertiserComponent;
  let fixture: ComponentFixture<CreateAdvertiserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAdvertiserComponent
      ],
      providers: [
        AdvertiserService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAdvertiserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});