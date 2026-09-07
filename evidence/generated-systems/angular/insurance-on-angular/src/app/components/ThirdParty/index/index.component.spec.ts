
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexThirdPartyComponent } from './index.component';
import { ThirdPartyService } from '../../../services/ThirdParty.service';

describe('IndexThirdPartyComponent', () => {
  let component: IndexThirdPartyComponent;
  let fixture: ComponentFixture<IndexThirdPartyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexThirdPartyComponent
      ],
      providers: [
        ThirdPartyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexThirdPartyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});