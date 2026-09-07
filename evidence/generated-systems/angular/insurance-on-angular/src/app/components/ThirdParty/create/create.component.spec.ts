
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateThirdPartyComponent } from './create.component';
import { ThirdPartyService } from '../../../services/ThirdParty.service';
import { Router } from '@angular/router';

describe('CreateThirdPartyComponent', () => {
  let component: CreateThirdPartyComponent;
  let fixture: ComponentFixture<CreateThirdPartyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateThirdPartyComponent
      ],
      providers: [
        ThirdPartyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateThirdPartyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});